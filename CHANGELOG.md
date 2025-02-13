# Change Log

All notable changes to this project will be documented in this file.
See [Conventional Commits](https://conventionalcommits.org) for commit guidelines.



# [0.3.7](https://github.com/tangx/goenv/compare/v0.3.6...v0.3.7)

### Bug Fixes

* **fix** 取消 fv.Set() 修改对象， 避免对象没有 SetDefaults() 和 Init() 而错误初始化引发的 panic ([acf5b01](https://github.com/tangx/goenv/commit/acf5b013316d75f52bdd1b00e666cdf6dc0ff536))



# [0.3.6](https://github.com/tangx/goenv/compare/v0.3.5...v0.3.6)

### Bug Fixes

* **fix** private filed no need to set ([b7f7c0f](https://github.com/tangx/goenv/commit/b7f7c0f8543d2e3838fc5ad73c3112bb6263d0b5))



# [0.3.5](https://github.com/tangx/goenv/compare/v0.3.4...v0.3.5)

### Bug Fixes

* **fix** 当字段对象类型为指针，且值为 nil 时， 使用反射进行初始化 ([a6192b8](https://github.com/tangx/goenv/commit/a6192b8b5ffaf4ffa9d0081b6dff4c8913649133))



# [0.3.4](https://github.com/tangx/goenv/compare/v0.3.3...v0.3.4)

### Bug Fixes

* **fix** fix default name with no tag env ([ce9f1f1](https://github.com/tangx/goenv/commit/ce9f1f1feb4c1a2790978cc2ad0bf11186dfcad5))



# [0.3.3](https://github.com/tangx/goenv/compare/v0.3.2...v0.3.3)

### Bug Fixes

* **fix** env - skip ([6ba3c68](https://github.com/tangx/goenv/commit/6ba3c6894d0279f57d3d88fa72d68be996af7cf7))



# [0.3.2](https://github.com/tangx/goenv/compare/v0.3.1...v0.3.2)

### Bug Fixes

* **fix** CallMethods 增加自定义方法调用 ([89fc7aa](https://github.com/tangx/goenv/commit/89fc7aadf53021de19288c59e900d6908544b773))



# [0.3.1](https://github.com/tangx/goenv/compare/v0.3.0...v0.3.1)

### Bug Fixes

* **fix** call SetDefaults before Init ([8c31483](https://github.com/tangx/goenv/commit/8c314832988f4ed228bc9a8482f1460404df3b1b))



# [0.3.0](https://github.com/tangx/goenv/compare/v0.2.1...v0.3.0)

### Features

* **feat** break change , renmae method caller ([923a833](https://github.com/tangx/goenv/commit/923a83399419a1229eea30581ccd521db42e5fbc))



# [0.2.1](https://github.com/tangx/goenv/compare/v0.2.0...v0.2.1)

### Bug Fixes

* **fix** lost value when trans env string slice into map ([ca10e1c](https://github.com/tangx/goenv/commit/ca10e1c057193283ef308ae708ef421de3d1ec1b))



# [0.2.0](https://github.com/tangx/goenv/compare/v0.1.0...v0.2.0)

### Features

* **feat** 增加读取文件 ([cbb5b90](https://github.com/tangx/goenv/commit/cbb5b901670a2ad9a0f36fea2af1b40884bea183))



# [0.1.0](https://github.com/tangx/goenv/compare/v0.0.1...v0.1.0)

### Features

* **feat** 使用结构体方法设置默认是 ([11c3dc1](https://github.com/tangx/goenv/commit/11c3dc16f8adaf57897d05d09754e3700ed01d5d))



# [0.0.1](https://github.com/tangx/goenv/compare/v0.0.0...v0.0.1)

### Bug Fixes

* **fix** load env error ([9db1980](https://github.com/tangx/goenv/commit/9db198087aa0a263f6140aa1891254f375f357d9))



# 0.0.0

### Bug Fixes

* **fix** skip unexported struct ([65b909f](https://github.com/tangx/goenv/commit/65b909f3890891b61441d497b673264a40570fe8))
* **fix** skip unexported struct ([09d5918](https://github.com/tangx/goenv/commit/09d59181b0575bb357d337d5a843ff094d391cd7))
* **fix** 修复了 field 判断优先级的问题 ([731ccf5](https://github.com/tangx/goenv/commit/731ccf5efb78e1f0e3e16331ce550a341ed647ce))


### Features

* **feat** marshal 支持 struct env tag ([147b2f0](https://github.com/tangx/goenv/commit/147b2f0fe73e3ecae95daa76e736495680585b80))
* **feat** unmarshal 支持 struct env tag ([00fe48b](https://github.com/tangx/goenv/commit/00fe48b718652894397e38747951f50edb0a75fd))
* **feat** 增加 unmarshal ([e3870ef](https://github.com/tangx/goenv/commit/e3870ef5013e25dea5c596702f535247e0b5b085))
