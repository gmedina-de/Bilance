package main

import (
	_ "genuine/app"
	"genuine/core"
	"genuine/core/server"
	"genuine/core/translator"
	"sync"
)

func main() {
	core.Invoke(func(server []server.Server, translator translator.Translator) any {
		var wg sync.WaitGroup
		for _, s := range server {
			s := s
			wg.Add(1)
			go func() {
				s.Serve()
				wg.Done()
			}()
		}
		wg.Wait()
		return nil
	})
}
