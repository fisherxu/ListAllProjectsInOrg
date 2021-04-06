## List all the projects in one github org

### Usage

```shell script
$ ListProjectsInOrg -h
Usage of ListProjectsInOrg:
  -org-name string
        The org name you want to list (default "kubernetes")
  -sort-by string
        The field name you want to sort by, include: star, fork, issue, createtime, updatetime, pushedtime. (default: star) (default "star")
```
                
### Example

```shell script
$ ListProjectsInOrg --org-name kubeedge --sort-by star
repos count:  10
┌────────────────────┬──────────┬───────┬───────┬───────────────────┬──────────────────────┬──────────────────────┬──────────────────────┬────────────────────────────────────────────────────────────────────────────────────┐
│ Name               │ Language │ Stars │ Forks │ Open_issues_count │ Created_at           │ Updated_at           │ Pushed_at            │ Description                                                                        │
├────────────────────┼──────────┼───────┼───────┼───────────────────┼──────────────────────┼──────────────────────┼──────────────────────┼────────────────────────────────────────────────────────────────────────────────────┤
│ kubeedge           │ Go       │ 3760  │ 983   │ 218               │ 2018-09-28T08:57:49Z │ 2021-04-06T08:09:51Z │ 2021-04-06T07:25:33Z │ Kubernetes Native Edge Computing Framework (project under CNCF)                    │
│ examples           │ Go       │ 118   │ 100   │ 15                │ 2018-12-19T08:05:12Z │ 2021-04-02T02:19:10Z │ 2021-03-17T06:16:31Z │ Examples for KubeEdge                                                              │
│ sedna              │ Go       │ 58    │ 17    │ 17                │ 2021-01-28T03:24:13Z │ 2021-04-06T08:45:15Z │ 2021-04-02T03:56:42Z │ AI tookit over KubeEdge                                                            │
│ website            │ HTML     │ 28    │ 56    │ 18                │ 2018-11-16T03:57:45Z │ 2021-04-01T11:36:29Z │ 2021-04-01T13:04:38Z │ KubeEdge website and  documentation repo:                                          │
│ beehive            │ Go       │ 26    │ 29    │ 0                 │ 2019-03-19T04:33:05Z │ 2021-03-12T14:25:45Z │ 2021-04-02T05:19:31Z │ Pluggable module(in-process microservice) and cross-module communication Framework │
│ viaduct            │ Go       │ 12    │ 16    │ 8                 │ 2019-03-25T06:46:20Z │ 2021-03-22T01:37:07Z │ 2020-11-30T06:40:12Z │ Wrapper for quic and websocket                                                     │
│ community          │          │ 11    │ 9     │ 5                 │ 2019-10-18T11:45:55Z │ 2021-03-02T15:33:57Z │ 2021-03-31T10:16:16Z │ KubeEdge community relevant content                                                │
│ mappers-go         │ Go       │ 7     │ 12    │ 5                 │ 2020-12-23T08:30:34Z │ 2021-02-27T11:25:28Z │ 2021-04-03T06:57:57Z │ KubeEdge Device Mappers written in go                                              │
│ katacoda-scenarios │          │ 5     │ 4     │ 0                 │ 2019-04-17T06:51:05Z │ 2019-12-22T13:45:27Z │ 2019-07-21T14:53:37Z │ Katacoda Interactive Scenarios                                                     │
│ community-infra    │          │ 0     │ 1     │ 3                 │ 2020-09-04T01:32:53Z │ 2021-01-07T06:44:01Z │ 2021-01-14T08:59:31Z │ Infrastructure for KubeEdge project                                                │
└────────────────────┴──────────┴───────┴───────┴───────────────────┴──────────────────────┴──────────────────────┴──────────────────────┴────────────────────────────────────────────────────────────────────────────────────┘
```