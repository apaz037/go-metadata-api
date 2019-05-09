# github.com/apaz037/go-metadata-api

A RESTful Golang API for persisting application metadata

## Routes

<details>
<summary>`/application`</summary>

- [RequestID]()
- [RealIP]()
- [Recoverer]()
- [Logger]()
- [Timeout.func1]()
- [SetContentType.func1]()
- [Heartbeat.func1]()
- **/application**
	- _GET_
		- [getAllApplications.func1](/api/handlers.go#L34)
	- _POST_
		- [createApplication.func1](/api/handlers.go#L16)

</details>
<details>
<summary>`/application/search/{params}`</summary>

- [RequestID]()
- [RealIP]()
- [Recoverer]()
- [Logger]()
- [Timeout.func1]()
- [SetContentType.func1]()
- [Heartbeat.func1]()
- **/application/search/{params}**
	- _GET_
		- [searchApplications.func1](/api/handlers.go#L40)

</details>
<details>
<summary>`/application/{id}`</summary>

- [RequestID]()
- [RealIP]()
- [Recoverer]()
- [Logger]()
- [Timeout.func1]()
- [SetContentType.func1]()
- [Heartbeat.func1]()
- **/application/{id}**
	- _DELETE_
		- [deleteApplication.func1](/api/handlers.go#L22)
	- _PUT_
		- [updateApplication.func1](/api/handlers.go#L28)
	- _GET_
		- [getApplication.func1](/api/handlers.go#L10)

</details>

Total # of routes: 3
