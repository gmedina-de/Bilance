package main

import (
	_ "genuine/app/accounting"
	_ "genuine/app/assets"
	_ "genuine/app/calendar"
	_ "genuine/app/common"
	_ "genuine/app/contacts"
	_ "genuine/app/dashboard"
	_ "genuine/app/files"
	_ "genuine/app/settings"
	"genuine/core"
	"genuine/core/server"
	"sync"
)

func main() {
	core.Invoke(func(server []server.Server) any {
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
