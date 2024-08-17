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

| Version      | Runtime of CodeGen | Changes                                 | Commit Hash                              |
| ------------ | ------------------ | --------------------------------------- | ---------------------------------------- |
| v0.0.1       | 1083661 ns/op      | Initial Version                         | 3cce7d76d78fd76fc7b63886077a6eb47caa61e6 |
| v0.0.2       | 536149 ns/op       | Format Builder in memory                | 96a9f0a46cee026e7476ff42978305f5a0f27af3 |
| v0.0.2 (alt) | 823289 ns/op       | Use templates instead of string builder | 340359aea5b3c8ab15a26282a4514581ae8d73b6 |
| v0.0.3       | 483838 ns/op       | Optimize keyword check                  |                                          |