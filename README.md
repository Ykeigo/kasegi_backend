# kasegi-backend

kasegi（仮）のバックエンド


デプロイ手順
1. EC2を立ち上げる（ここまではハンズオンの資料でなんとかして）
2. ECRにログインする
sudo aws ecr get-login-password --region ap-northeast-1 | sudo docker login --username AWS --password-stdin 744778349286.dkr.ecr.ap-northeast-1.amazonaws.com
3. EC2にログインしてECRからpullする

ssh -o ProxyCommand='ssh -i oyanagi_practice_app_key.pem -W %h:%p ec2-user@ec2-52-195-5-241.ap-northeast-1.compute.amazonaws.com' -i oyanagi_practice_app_key.pem ec2-user@10.0.134.70

aws ecr get-login-password --region ap-northeast-1 | docker login --username AWS --password-stdin 461711549202.dkr.ecr.ap-northeast-1.amazonaws.com
sudo aws ecr get-login-password --region ap-northeast-1 | sudo docker login --username AWS --password-stdin 461711549202.dkr.ecr.ap-northeast-1.amazonaws.com