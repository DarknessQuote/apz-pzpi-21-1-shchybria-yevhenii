import { useTranslation } from "react-i18next";
import { useAuthContext } from "../context/AuthContext";
import { Typography } from "@mui/material";
import { useEffect, useState } from "react";
import { getUser } from "../services/userService";
import UserInfo from "../components/UserInfo";

const HomePage = () => {
	const [auth] = useAuthContext();
	const [user, setUser] = useState(null);

	const { t } = useTranslation();

	useEffect(() => {
		const getUserInfo = async () => {
			if (auth !== null) {
				const userData = await getUser(auth.userID);
				setUser(userData);
			}
		};

		getUserInfo();
	}, [auth]);

	if (auth == null) {
		return <Typography>{t("noAuth")}</Typography>;
	} else if (user !== null) {
		return <UserInfo userInfo={user} />;
	}
};

export default HomePage;
