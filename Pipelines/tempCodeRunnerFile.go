go func() {
		wg.Wait()
		close(results)
	}()