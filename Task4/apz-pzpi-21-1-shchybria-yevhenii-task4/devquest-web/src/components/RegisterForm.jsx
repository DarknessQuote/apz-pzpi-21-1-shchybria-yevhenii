import {
	Box,
	Button,
	FormControl,
	InputLabel,
	MenuItem,
	Select,
	TextField,
} from "@mui/material";
import { useEffect, useRef, useState } from "react";
import { getRoles } from "../services/authService.js";
import { getCompanies } from "../services/companyService.js";

const RegisterForm = ({ authenticateUser }) => {
	const [roles, setRoles] = useState([]);
	const [companies, setCompanies] = useState([]);

	const usernameRef = useRef();
	const firstNameRef = useRef();
	const lastNameRef = useRef();
	const passwordRef = useRef();
	const roleRef = useRef();
	const companyRef = useRef();

	useEffect(() => {
		const getDataForRegister = async () => {
			const roles = await getRoles();
			setRoles(roles);

			const companies = await getCompanies();
			setCompanies(companies);
		};

		getDataForRegister();
	}, []);

	const handleSubmit = async (e) => {
		e.preventDefault();

		const registerData = {
			username: usernameRef.current.value,
			firstName: firstNameRef.current.value,
			lastName: lastNameRef.current.value,
			password: passwordRef.current.value,
			role: roleRef.current.value,
			company: companyRef.current.value,
		};

		await authenticateUser(registerData);
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
					name="firstName"
					label="First Name"
					variant="outlined"
					inputRef={firstNameRef}
					InputLabelProps={{ shrink: true }}
				/>
				<TextField
					required
					name="lastName"
					label="Last Name"
					variant="outlined"
					inputRef={lastNameRef}
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
				<FormControl>
					<InputLabel>Role</InputLabel>
					<Select
						required
						title="role"
						label="Role"
						inputRef={roleRef}>
						<MenuItem selected value="">
							Select role
						</MenuItem>
						{roles.map((role) => {
							return (
								<MenuItem value={role.id} key={role.id}>
									{role.title}
								</MenuItem>
							);
						})}
					</Select>
				</FormControl>
				<FormControl>
					<InputLabel>Company</InputLabel>
					<Select
						required
						title="company"
						label="Company"
						inputRef={companyRef}>
						<MenuItem selected value="">
							Select company
						</MenuItem>
						{companies.map((company) => {
							return (
								<MenuItem value={company.id} key={company.id}>
									{company.name}
								</MenuItem>
							);
						})}
					</Select>
				</FormControl>
				<Button type="submit">Register</Button>
			</Box>
		</form>
	);
};

export default RegisterForm;
