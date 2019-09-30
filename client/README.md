# procon-qa

## Project setup
```
npm install
```

### Compiles and hot-reloads for development
```
npm run serve
```

```
npm ERR! code ELIFECYCLE
npm ERR! syscall spawn
npm ERR! file sh
npm ERR! errno ENOENT
npm ERR! procon-qa@0.1.0 build: `vue-cli-service build`
npm ERR! spawn ENOENT
npm ERR! 
npm ERR! Failed at the procon-qa@0.1.0 build script.
npm ERR! This is probably not a problem with npm. There is likely additional logging output above.
npm WARN Local package.json exists, but node_modules missing, did you mean to install?

npm ERR! A complete log of this run can be found in:
npm ERR!     /Users/monkukui/.npm/_logs/2019-09-17T11_13_26_064Z-debug.log
```
上のようなエラーが出た場合, 以下のコマンドを実行する
```
rm -rf node_modules
npm install -g npm@latest
npm i core-util-is
```

### Compiles and minifies for production
```
npm run build
```
/dist 以下に *html, *js, *css ファイルが生成される

### Run your tests
```
npm run test
```

### Lints and fixes files
```
npm run lint
```
commit する前は必ず実行する

### Customize configuration
See [Configuration Reference](https://cli.vuejs.org/config/).
