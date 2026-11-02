# llm-gateway: seven-day MVP

Goal: Route model requests through a consistent backend interface.

Status: queued. The initial core component is implemented; the full MVP is not.

- [ ] Day 1: Provider selection policy and tests.
- [ ] Day 2: Define request/response contract and mock provider.
- [ ] Day 3: Add timeouts and bounded retries.
- [ ] Day 4: Implement rate limiting and usage accounting.
- [ ] Day 5: Add safe cache policy for eligible requests.
- [ ] Day 6: Test provider failures and routing behavior.
- [ ] Day 7: Publish local gateway demo and benchmarks.

## Definition of done

- A documented local command runs the scoped end-to-end example.
- Meaningful positive and failure-path tests pass locally and in CI.
- README distinguishes implemented features, mocks, measured results, and limitations.
- Example inputs and outputs are reproducible.
- A short demo script and evidence-based LinkedIn draft are saved locally.
- Real model quality is not claimed without a live evaluation.

Seven days is a target. Reduce scope or record unfinished work explicitly; never mark a project complete merely because its week ended.
