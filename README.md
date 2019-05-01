# PoC WebApp for APPUiO Admin API

Consumes API from [cdays-webapi-poc](https://github.com/vshn/cdays-webapi-poc).

## Development

With [refresh](https://github.com/markbates/refresh):

```
refresh -c refresh.yml
```

## Swagger Client

Generate the Swagger client using:

```
swagger generate client -f ../cdays-webapi-poc/swagger.yaml -A webapi
```