* release notes:
* 0.12.0-rc1 removes plenty of technical debt and includes many bug-fixes. Several user-centric upgrades are also included (read on :doug:)
* the most significant user-facing update is the switch from [ubuntu](https://github.com/eris-ltd/common/blob/875c64ec7e1525b3c8e0efed7d4974665fff4942/docker/base-ubuntu/Dockerfile) to [alpine](https://github.com/eris-ltd/common/blob/875c64ec7e1525b3c8e0efed7d4974665fff4942/docker/base-alpine/Dockerfile) as the base image in eris' base docker image. This included the addition of a [build](https://github.com/eris-ltd/common/blob/875c64ec7e1525b3c8e0efed7d4974665fff4942/docker/build/Dockerfile) image from which each of eris-keys, eris-cm, eris-db, and eris-pm are now based from. Concurrently, the keys (#728) and data (#733) images are now versioned - providing additional stability and flexibility moving forward. In sum, `eris init`now pulls in ~400MB rather than ~1G of docker imagery. Doug lost weight.
* the smaller sized images also benefits the portability of running eris on IoT devices. This release introduces native support for ARM based architechture (#719). See the [common Dockerfiles here](https://github.com/eris-ltd/common/tree/875c64ec7e1525b3c8e0efed7d4974665fff4942/docker-arm) and the `Dockerfile.armhf` in each the eris-keys, eris-cm, eris-db, and eris-pm repositories.
* the `pkgs` command has a fixed directory check (#688, #723)
* the `pkgs` command now has `import/export` functionality (#689) similar to `services/actions`
* the `chains graduate` command has been deprecated (#690)
* the `eth` service has been renamed to `geth` to better reflect the service (#691)
* support for (native) Docker on OSX/Windows added (#693)
* the `--api` flag on ` chains new/start` has been removed (#694). The api will always be on.
* cleaner error handling added to the agent package (#695)
* `eris init` displays the directory structure at the end with additional helpful information (#708)
* **deprecation notice:** `services new` is no longer and has been renamed `services make` (#713)
* informative commands (`version/update/man`) no longer connect to docker since they don't need to (#716)
* a new [external library "eris-logger"](https://github.com/eris-ltd/eris-logger) replaces the use of the logrus library (#720)
* tests on CircleCI now run in parallel, reducing test time (#736)
* loaders package tests were cleaned up & score on goreportcard improved (#747)
* fixes to `clean` and image removal in general were added, along with built-on-the-fly images for testing (#755)

* eris-keys:
* https://github.com/eris-ltd/eris-keys/pull/55

* eris-cm:

* eris-db:

* eris-pm:
