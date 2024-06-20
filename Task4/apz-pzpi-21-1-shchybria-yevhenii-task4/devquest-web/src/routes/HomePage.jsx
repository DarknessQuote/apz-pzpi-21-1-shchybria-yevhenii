import { useAuthContext } from "../context/AuthContext";
import { Typography } from "@mui/material";

const HomePage = () => {
	const [auth] = useAuthContext();

	if (auth == null) {
		return <Typography>Not authenticated</Typography>;
	} else {
		return (
			<>
				<Typography overflow="clip">Token: {auth.token}</Typography>
				<Typography>ID: {auth.userID}</Typography>
				<Typography>Role: {auth.role}</Typography>
			</>
		);
	}
};

export default HomePage;
