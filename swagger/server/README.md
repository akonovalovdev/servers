Написанной от руки спецификацией OpenAPI является task.yaml.

Серверный код генерируется с помощью:

java -jar swagger-codegen-cli.jar generate -i task.yaml -l go-server -o ./swag-server
Где swagger-codegen-cli.jar это загруженный двоичный файл (для OpenAPI версии 3) с веб-сайта Swagger Codegen.

Обратите внимание, что при этом также создается копия спецификации в ./swag-server/api/swagger.yaml.
Эта копия немного отличается от входной task.yaml - она канонизирована и переформатирована генератором кода.

Пришлось переименовать go dir в swagger, чтобы привести в соответствие с именем пакета (для модулей).

task-swagger-2.json  это наш API, преобразованный в Swagger (OpenAPI v2) с помощью онлайн-инструмента  https://lucybot-inc.github.io/api-spec-converter/

oapi-server основан на коде, сгенерированном с помощью инструмента deepmap/oapi-codegen tool:

oapi-codegen -package task -generate types,server task.yaml > oapi-server/internal/task/task.gen.go
Файл .gen.go не изменен, за исключением добавления обработчика "удалить все", который не является частью общедоступного API.