# Teval

Template evaluator for working with different manifests in this repository.

For the most part, adopting an application to this project is:

1. Wrapping the container with Litestream.
1. Deploy with ~80% same YAML lines.

## Alternatives

I considered several options before finally deciding to handroll (heh).

### YTT

Honestly quite neat and was going to be the top choice, but the syntax has some learning curve and I'm not sure if I have time from my weekends to do that.

### Helm

Powerful tool, but after working with several charts over the years, I never really liked the debugging experience. I want to see literal YAML when things go wrong.


### TrueNAS / TrueCharts

I like the idea and applaud how much they manage to offer with any tooling they have, despite still using Helm under the hood. Unfortunately, maintainers seem to limit options strictly, such as only supporting PostgreSQL whenever apps offer them.

I would use PostgreSQL too when I expect significant traffic. However, this is my home cluster, which has average user count of 1.5 people per application. Managing N databases for N applications doesn't feel like a good use of my time.
