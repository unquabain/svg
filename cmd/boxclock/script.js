function display(unit, multiplier, max, value) {
	const idPrefix = '' + unit + '-' + multiplier + '-'
	for (var i = max; i > 0; i -= multiplier) {
			const id = idPrefix + i
			if (value >= i) {
				document.getElementById(id).style.display = 'block'
			} else {
				document.getElementById(id).style.display = 'none'
			}
	}
}
setInterval(function() {
	const now = new Date()
	const t = Math.floor(now.getTime() / 1000 - now.getTimezoneOffset()*60)
	const s = t % 60
	const m = Math.floor(t / 60) % 60 
	const h = Math.floor(t / 3600) % 24

	display('s', 1, 4, s % 5)
	display('s', 5, 10, s % 15)
	display('s', 15, 45, s % 60)

	display('m', 1, 4, m % 5)
	display('m', 5, 10, m % 15)
	display('m', 15, 45, m % 60)

	display('h', 1, 2, h % 3)
	display('h', 3, 9, h % 12)
	display('h', 12, 12, h % 24)
}, 500)
