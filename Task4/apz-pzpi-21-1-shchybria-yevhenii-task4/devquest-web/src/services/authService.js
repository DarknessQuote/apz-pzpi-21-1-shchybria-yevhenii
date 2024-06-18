export const register = async (registerData) => {
	try {
		const headers = new Headers();
		headers.append("Content-Type", "application/json");

		const reqBody = {
			username: registerData.username,
			first_name: registerData.firstName,
			last_name: registerData.lastName,
			password: registerData.password,
			role_id: registerData.roleID,
			company_id: registerData.companyID,
		};

		const reqOptions = {
			method: "POST",
			body: JSON.stringify(reqBody),
			headers: headers,
		};

		const responseJSON = await fetch(
			"http://localhost:8080/auth/register",
			reqOptions
		);
		const response = await responseJSON.json();

		if (response.error) {
			throw new Error(response.message);
		}

		return response.data;
	} catch (err) {
		console.error(err);
	}
};

export const login = async (loginData) => {
	try {
		const headers = new Headers();
		headers.append("Content-Type", "application/json");

		const reqBody = {
			username: loginData.username,
			password: loginData.password,
		};

		const reqOptions = {
			method: "POST",
			body: JSON.stringify(reqBody),
			headers: headers,
		};

		const responseJSON = await fetch(
			"http://127.0.0.1:8080/auth/login",
			reqOptions
		);
		const response = await responseJSON.json();

		if (response.error) {
			throw new Error(response.message);
		}

		return response.data;
	} catch (err) {
		console.error(err);
	}
};

export const logout = async () => {
	try {
		const reqOptions = {
			method: "DELETE",
			credentials: "include",
		};

		await fetch("http://127.0.0.1:8080/auth/logout", reqOptions);
	} catch (err) {
		console.log(err);
	}
};
