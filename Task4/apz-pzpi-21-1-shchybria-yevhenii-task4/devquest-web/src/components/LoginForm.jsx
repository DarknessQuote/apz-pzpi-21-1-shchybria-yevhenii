import { Box, Button, TextField } from "@mui/material";
import { useRef } from "react";

const LoginForm = ({ authenticateUser }) => {
	const usernameRef = useRef();
	const passwordRef = useRef();

	const handleSubmit = async (e) => {
		e.preventDefault();

		const loginData = {
			username: usernameRef.current.value,
			password: passwordRef.current.value,
		};

		await authenticateUser(loginData);
	};

	return (
		<form onSubmit={handleSubmit}>
			<Box className="flex flex-col gap-3">
				<TextField
					required
					name="username"
					label="Username"
					variant="outlined"
					inputRef={usernameRef}
					InputLabelProps={{ shrink: true }}
				/>
				<TextField
					required
					name="password"
					type="password"
					label="Password"
					variant="outlined"
					inputRef={passwordRef}
					InputLabelProps={{ shrink: true }}
				/>
				<Button type="submit">Login</Button>
			</Box>
		</form>
	);
};

export default LoginForm;
