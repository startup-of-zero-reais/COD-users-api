read -p "Tecle enter para enviar sua .env.prod para produção: " o

ENV_LOCATION=$(pwd)
APPLICATION_S3_FOLDER="cod-users-api"

function send_file() {
	aws=$(aws --version)
    if [[ $aws == "" ]]; then
    	echo "Voce deve instalar a CLI 2.0 da AWS"
    	echo "https://docs.aws.amazon.com/pt_br/cli/latest/userguide/install-cliv2.html"
	else
		aws s3api put-object --bucket code-craft-envs --key $APPLICATION_S3_FOLDER/.env --body $ENV_LOCATION/$ENV_FILE
	fi
}

ENV_FILE=$o

function env_file() {
	if [[ $o != "" ]]; then
		ENV_FILE=".env.prod"
		echo "Iniciando envio de: .env.prod..."
	else
		ENV_FILE=$o
		echo "Iniciando envio de: $o..."
	fi
}

function search_file() {
    if [ -f "$ENV_LOCATION/$ENV_FILE" ]; then
    	send_file
	else
    	echo "$ENV_LOCATION/$ENV_FILE - nao encontrado"
	fi
}

env_file
search_file