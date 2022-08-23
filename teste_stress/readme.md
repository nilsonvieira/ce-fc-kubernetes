# Versão da imagem para o teste de stress
Na aula "Teste de stress com fortio" a imagem que está sendo utilizada é a wesleywillians/hello-go:v9.6 que é diferente da versão v5.5 usada no deployment.yaml da aula "Aplicando deployment com resources".
 
O que muda é que na versão 5.5 o código do server.go gerava erro após 30 segundos de execução no pod.

Na versão 9.6 este erro após 30 segundos é removido e o teste de stress funcionará normalmente.
 
Versão 5.5 (wesleywillians/hello-go:5.5)
 
```go
func Healthz(w http.ResponseWriter, r *http.Request) {
  duration := time.Since(startedAt)
 
  if duration.Seconds() < 10 || duration.Seconds() > 30 {
    w.WriteHeader(500)
    w.Write([]byte(fmt.Sprintf("Duration: %v", duration.Seconds())))
  } else {
    w.WriteHeader(200)
    w.Write([]byte("ok"))
  }
}
 
Na versão 9.6 (wesleywillians/hello-go:v9.6)
 
func Healthz(w http.ResponseWriter, r *http.Request) {
  duration := time.Since(startedAt)
 
  if duration.Seconds() < 10 {
    w.WriteHeader(500)
    w.Write([]byte(fmt.Sprintf("Duration: %v", duration.Seconds())))
  } else {
    w.WriteHeader(200)
    w.Write([]byte("ok"))
  }
}
```
Portanto, para que você consiga acompanhar a aula "Teste de stress com fortio", utilize a
imagem wesleywillians/hello-go:v9.6 no deployment.yaml.

# Atualização no comando do Fortio
Olá pessoal, tudo bem?

Com a atualização do kubectl para a versão 1.21, o parâmetro `--generator` do comando abaixo passou a não ser mais suportado, apresentando o erro **"Error: unknown flag: --generator".**

```bash
$ kubectl run -it --generator=run-pod/v1 fortio --rm --image=fortio/fortio -- load -qps 800 -t 120s -c 70 "http://goserver-service/healthz"
```



Para a realização do teste, execute o comando sem este parâmetro, ficando da seguinte maneira:
```bash
$ kubectl run -it fortio --rm --image=fortio/fortio -- load -qps 800 -t 120s -c 70 "http://goserver-service/healthz"
```

Com isto, será possível ver os pods escalando no teste de stress.

É isto aí e até a próxima!