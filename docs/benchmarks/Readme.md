# Benchmarks

In this page, we will be showcasing the different benchmarks that we have conducted to see how the system has improved over the different commits, including their changes.

For each of the releases, I will aim to make this builder faster.

## Testing hardware

This is completed on the following hardware
1. CPU: AMD Ryzen 5 7600 (6 core processor)
2. RAM: 32GB at 6000MT/s
3. Storage: Samsung 960 Evo 500GB

## Testing methodology

The test is benchmark based on the time it takes to generate the builder and write the result to disk.
For commits that do not write to disk, we will be writing a disk writer file to test the disk write speed.

## Results

The old code only contains CodeGen + I/O mode as I/O is tightly coupled to the implementation.

| Version      | Runtime of CodeGen + I/O | Runtime of CodeGen only | Changes                                 |
| ------------ | ------------------------ | ----------------------- | --------------------------------------- |
| v0.0.1       | 1083661 ns/op            | -                       | Initial Version                         |
| v0.0.2       | 536149 ns/op             | -                       | Format Builder in memory                |
| v0.0.2 (alt) | 823289 ns/op             | -                       | Use templates instead of string builder |
| v0.0.3       | 483838 ns/op             | 267546 ns/op            | Optimize keyword check                  |
| v0.0.3 (alt) | 603091 ns/op             | 267859 ns/op            | Optimize keyword differently            |
| v0.0.4       | 293983 ns/op             | 51133 ns/op             | Manual format/import pkgs               |
| v0.0.5       | 418557 ns/op             | 202906 ns/op            | Fix import part using parser            |
| v0.0.6       | 283447 ns/op             | 45850 ns/op             | Update import part using parser         |
| v0.0.7       | 195134 ns/op             | 48860 ns/op             | Optimize string opts                    |
| v0.1.0       | 273165 ns/op             | 28255 ns/op             | Custom Parser                           |
| v0.2.0       | 282231 ns/op             | 30701 ns/op             | Parsing by file + fix errors            |
