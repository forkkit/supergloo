// Code generated by solo-kit. DO NOT EDIT.

package v1

import (
	"context"
	"fmt"
	"time"

	gloo_solo_io "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	github_com_solo_io_solo_kit_pkg_api_v1_resources_common_kubernetes "github.com/solo-io/solo-kit/pkg/api/v1/resources/common/kubernetes"

	"go.opencensus.io/stats"

	"github.com/solo-io/go-utils/errutils"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
)

type DiscoverySimpleEmitter interface {
	Snapshots(ctx context.Context) (<-chan *DiscoverySnapshot, <-chan error, error)
}

func NewDiscoverySimpleEmitter(aggregatedWatch clients.ResourceWatch) DiscoverySimpleEmitter {
	return NewDiscoverySimpleEmitterWithEmit(aggregatedWatch, make(chan struct{}))
}

func NewDiscoverySimpleEmitterWithEmit(aggregatedWatch clients.ResourceWatch, emit <-chan struct{}) DiscoverySimpleEmitter {
	return &discoverySimpleEmitter{
		aggregatedWatch: aggregatedWatch,
		forceEmit:       emit,
	}
}

type discoverySimpleEmitter struct {
	forceEmit       <-chan struct{}
	aggregatedWatch clients.ResourceWatch
}

func (c *discoverySimpleEmitter) Snapshots(ctx context.Context) (<-chan *DiscoverySnapshot, <-chan error, error) {
	snapshots := make(chan *DiscoverySnapshot)
	errs := make(chan error)

	untyped, watchErrs, err := c.aggregatedWatch(ctx)
	if err != nil {
		return nil, nil, err
	}

	go errutils.AggregateErrs(ctx, errs, watchErrs, "discovery-emitter")

	go func() {
		originalSnapshot := DiscoverySnapshot{}
		currentSnapshot := originalSnapshot.Clone()
		timer := time.NewTicker(time.Second * 1)
		sync := func() {
			if originalSnapshot.Hash() == currentSnapshot.Hash() {
				return
			}

			stats.Record(ctx, mDiscoverySnapshotOut.M(1))
			originalSnapshot = currentSnapshot.Clone()
			sentSnapshot := currentSnapshot.Clone()
			snapshots <- &sentSnapshot
		}

		defer func() {
			close(snapshots)
			close(errs)
		}()

		for {
			record := func() { stats.Record(ctx, mDiscoverySnapshotIn.M(1)) }

			select {
			case <-timer.C:
				sync()
			case <-ctx.Done():
				return
			case <-c.forceEmit:
				sentSnapshot := currentSnapshot.Clone()
				snapshots <- &sentSnapshot
			case untypedList := <-untyped:
				record()

				currentSnapshot = DiscoverySnapshot{}
				for _, res := range untypedList {
					switch typed := res.(type) {
					case *github_com_solo_io_solo_kit_pkg_api_v1_resources_common_kubernetes.Pod:
						currentSnapshot.Pods = append(currentSnapshot.Pods, typed)
					case *gloo_solo_io.Upstream:
						currentSnapshot.Upstreams = append(currentSnapshot.Upstreams, typed)
					case *TlsSecret:
						currentSnapshot.Tlssecrets = append(currentSnapshot.Tlssecrets, typed)
					case *github_com_solo_io_solo_kit_pkg_api_v1_resources_common_kubernetes.Deployment:
						currentSnapshot.Deployments = append(currentSnapshot.Deployments, typed)
					default:
						select {
						case errs <- fmt.Errorf("DiscoverySnapshotEmitter "+
							"cannot process resource %v of type %T", res.GetMetadata().Ref(), res):
						case <-ctx.Done():
							return
						}
					}
				}

			}
		}
	}()
	return snapshots, errs, nil
}
