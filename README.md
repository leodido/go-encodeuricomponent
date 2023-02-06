# go-encodeuricomponent

A little library that replicates the behavior of the `encodeURIComponent` JavaScript function.

##### Notice

The current implementation doesn't detect lead and trail surrogates.

The Javascript `encodeURIComponent` returns an error when one attempts to encode a lone low/high surrogate.

This library doesn't, at the moment. Supporting this feature is a very valuable way to start contributing!

Get to know more by reading the [mdn web docs](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/encodeURI#encoding_a_lone_high_surrogate_throws).
