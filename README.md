# Go Bindings & Example Project

For [Gamercade](https://gamercade.io) and the [Gamercade Console](https://github.com/gamercade-io/gamercade_console).

## Tools Used
Go - 1.18
TinyGo - 0.25.0
Gamercade - 0.1

## Go Bindings for Gamercade

Bindings are located at `/gamercade`.

Import the bindings into your project with `import "github.com/zyellowhorse/go-template/gamercade"`.

## Building the Example project

I used [TinyGo](https://tinygo.org/) to compile to WASM. You can install TinyGo by following the instructions found [here](https://tinygo.org/getting-started/install/)

Once installed compile to WASM with the following command 
`GOOS=js GOARCH=wasm tinygo build -o main.wasm`

From here you can use the Gamercade Editor to bundle with any assets to create create your game.

## License

Licensed under either of

 * Apache License, Version 2.0 ([LICENSE-APACHE](LICENSE-APACHE) or https://www.apache.org/licenses/LICENSE-2.0)
 * MIT license ([LICENSE-MIT](LICENSE-MIT) or https://opensource.org/licenses/MIT)


#### Contribution

Unless you explicitly state otherwise, any contribution intentionally submitted
for inclusion in the work by you, as defined in the Apache-2.0 license, shall be
dual licensed as above, without any additional terms or conditions.
