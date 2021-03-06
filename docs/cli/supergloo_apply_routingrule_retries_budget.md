---
title: "supergloo apply routingrule retries budget"
weight: 5
---
## supergloo apply routingrule retries budget

b

### Synopsis

b

```
supergloo apply routingrule retries budget [flags]
```

### Options

```
  -h, --help                 help for budget
  -m, --min-retries uint32   the proxy may always attempt this number of retries per second, even if it would violate the retryRatio
  -r, --ratio float32        the ratio of additional traffic that may be added by retries. retry_ratio of 0.1 means that 1 retry may be attempted for every 10 regular requests
  -t, --ttl duration         This duration indicates for how long requests should be considered for the purposes of enforcing the retryRatio.  A higher value considers a larger window and therefore allows burstier retries.
```

### Options inherited from parent commands

```
      --dest-labels MapStringStringValue       apply this rule to requests sent to pods with these labels. format must be KEY=VALUE (default [])
      --dest-namespaces strings                apply this rule to requests sent to pods in these namespaces
      --dest-upstreams ResourceRefsValue       apply this rule to requests sent to these upstreams. format must be <NAMESPACE>.<NAME>. (default [])
      --dryrun                                 if true, this command will print the yaml used to create a kubernetes resource rather than directly trying to create/apply the resource
  -i, --interactive                            run in interactive mode
      --name string                            name for the resource
      --namespace string                       namespace for the resource (default "supergloo-system")
  -o, --output string                          output format: (yaml, json, table)
      --request-matcher RequestMatchersValue   json-formatted string which can be parsed as a RequestMatcher type, e.g. {"path_prefix":"/users","path_exact":"","path_regex":"","methods":["GET"],"header_matchers":{"x-custom-header":"bar"}} (default [])
      --source-labels MapStringStringValue     apply this rule to requests originating from pods with these labels. format must be KEY=VALUE (default [])
      --source-namespaces strings              apply this rule to requests originating from pods in these namespaces
      --source-upstreams ResourceRefsValue     apply this rule to requests originating from these upstreams. format must be <NAMESPACE>.<NAME>. (default [])
      --target-mesh ResourceRefValue           select the target mesh or mesh group to which to apply this rule. format must be NAMESPACE.NAME (default { })
```

### SEE ALSO

* [supergloo apply routingrule retries](../supergloo_apply_routingrule_retries)	 - rt

