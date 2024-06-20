import {
	AppBar,
	Box,
	Button,
	ButtonGroup,
	Toolbar,
	Typography,
} from "@mui/material";
import { useAuthContext } from "../context/AuthContext";
import { NavLink, useNavigate } from "react-router-dom";
import { logout } from "../services/authService";

const Header = () => {
	const [auth, setAuth] = useAuthContext();
	const navigate = useNavigate();

	return (
		<AppBar position="sticky" className="mb-5">
			<Toolbar className="flex items-center">
				<Box className="grow">
					<Typography>DevQuest</Typography>
				</Box>
				{auth === null ? (
					<ButtonGroup variant="contained" disableElevation>
						<NavLink to="/auth?mode=login">
							<Button>Login</Button>
						</NavLink>
						<NavLink to="/auth?mode=register">
							<Button>Register</Button>
						</NavLink>
					</ButtonGroup>
				) : (
					<Button
						variant="contained"
						onClick={async () => {
							await logout();
							setAuth(null);
							navigate("/");
						}}>
						Log out
					</Button>
				)}
			</Toolbar>
		</AppBar>
	);
};

export default Header;
