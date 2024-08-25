compile:
	npx hardhat compile

deploy:
	npx hardhat ignition deploy ./ignition/modules/News.ts --network ganache

server:
	nodemon --watch 'frontend/**/*.html' --exec 'ts-node' frontend/server.ts
