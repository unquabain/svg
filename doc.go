/* svg: An SVG DOM building library.

The SVG package is little more than a set of structs with tagged properties.
When marshalled with the Go standard library's encoding/xml.Marshal or
encoding/xml.MarshalIndent, they should render as SVG.

The idea is not that it's easier to write Go than to write SVG, but rather
to make it easier to programmatically build SVG for some Go consumer, such as
fyne-io. The idea is that you could write a function or object that uses
these SVG primitives to produce an image that could be used as a part of a
UI: the background of a widget or some graphical control.

*/
package svg
