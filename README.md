# Handroll

Personal for self-hosting applications. This is an attempt to manage multiple deployments, variety of applications, while maintaining consistency and proper backups. Through the research journey, my conclusion and strategy boils down to:

- Kubernetes is convenient, especially if apps can be treated as "stateless".
- Consolidating states into a single protocol means I can use single backup strategy. Object storage makes a lot of sense to tackle this.
- As such, I opted to run S3-compatible server on node OS directly. I use [Garage](https://garagehq.deuxfleurs.fr) at the time of writing.
- The last thing is databases, typically SQL. I opted for always using SQLite, wrapping upstream images with [Litestream](https://litestream.io) for backups, thus removing the need for Persistent Volume Claims in most cases.

## Contributing & usage

If you would like to offer ideas where you can make use of this project, please open an issue and we can discuss what we can do.

I would encourage you to take a look, try things out, and (most importantly) make it your own!

