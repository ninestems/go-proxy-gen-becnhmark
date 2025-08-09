## Benchmark Results

The benchmarks were run on an Apple M1 Pro (ARM64) machine using:

```bash
go test -bench=. -benchmem -benchtime=10s -count=3 ./benchmark/
```

| Benchmark                 | Iterations    | Time per op (ns) | Memory per op (B) | Allocs per op |
| ------------------------- | ------------- | ---------------- | ----------------- | ------------- |
| `BenchmarkConverter_Gen`  | \~7,200,000   | \~1,660          | 1,803             | 15            |
| `BenchmarkConverter_Hand` | \~8,100,000   | \~1,490          | 1,730             | 10            |
| `BenchmarkConverter_Pure` | \~800,000,000 | \~15             | 0                 | 0             |

## Explanation of Implementations
- Gen — This is the generated converter implementation that uses reflection internally. 
It incurs higher memory allocations and is slower due to the overhead of reflection.
- Hand — This is a generated implementation without reflection, generated one-to-one, 
resulting in fewer allocations and better performance compared to the Gen version.
- Pure — A bare, direct call to the implementation without any wrappers or reflection. 
It is extremely fast and allocates no memory.

## Extra
For logger and tracer, no-op implementations were used to avoid affecting benchmark results.