import { useTranslation } from "react-i18next";
import { useAuthContext } from "../context/AuthContext";
import { Typography } from "@mui/material";

const HomePage = () => {
	const [auth] = useAuthContext();

	const { t } = useTranslation();

	if (auth == null) {
		return <Typography>{t("noAuth")}</Typography>;
	} else {
		return (
			<>
				<Typography className="break-words">
					Token: {auth.token}
				</Typography>
				<Typography>ID: {auth.userID}</Typography>
				<Typography>Role: {auth.role}</Typography>
			</>
		);
	}
};

export default HomePage;
