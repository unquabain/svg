const go = new globalThis.Go()
WebAssembly.instantiate(Uint8Array.from(atob(%q), c => c.charCodeAt(0)),go.importObject)
  .then(result => {
    go.run(result.instance)
    setInterval(tick, 500)
  })
