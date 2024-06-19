import { Link } from "react-router-dom";
import { useAuthContext } from "../context/AuthContext";
import { Button } from "@mui/material";
import { logout } from "../services/authService";

const HomePage = () => {
	const [auth, setAuth] = useAuthContext();

	if (auth == null) {
		return (
			<>
				<Link to="/auth?mode=login">Login</Link>
				<br />
				<Link to="/auth?mode=register">Register</Link>
			</>
		);
	} else {
		return (
			<>
				<p>Token: {auth.token}</p>
				<p>ID: {auth.userID}</p>
				<p>Role: {auth.role}</p>
				<Button
					onClick={async () => {
						await logout();
						setAuth(null);
					}}>
					Logout
				</Button>
			</>
		);
	}
};

export default HomePage;
