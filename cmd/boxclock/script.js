const go = new globalThis.Go()
WebAssembly.instantiate(Uint8Array.from([%v]),go.importObject)
  .then(result => {
    go.run(result.instance)
    setInterval(tick, 500)
  })
