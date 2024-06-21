import { useEffect, useState } from "react";
import { deleteCompany, getCompanies } from "../services/companyService";
import {
	Button,
	ButtonGroup,
	Divider,
	List,
	ListItem,
	ListItemButton,
	ListItemText,
	Paper,
} from "@mui/material";
import { useTranslation } from "react-i18next";
import { useAuthContext } from "../context/AuthContext";

const CompaniesPage = () => {
	const [companies, setCompanies] = useState([]);

	const [auth] = useAuthContext();

	const { t } = useTranslation();

	useEffect(() => {
		const loadCompanies = async () => {
			const companiesList = await getCompanies();
			setCompanies(companiesList);
		};

		loadCompanies();
	}, []);

	const handleDeleteCompany = async (companyID) => {
		await deleteCompany(companyID, auth.token);
		setCompanies(await getCompanies());
	};

	return (
		<>
			<Paper>
				<List>
					{companies.map((company, i) => {
						return (
							<>
								<ListItem
									className="flex justify-start"
									key={company.id}>
									<ListItemText
										primary={company.name}
										secondary={company.owner}
										className="w-96 grow-0"
									/>
									<ListItemText
										secondary={company.email}
										className="grow"
									/>
									<ButtonGroup variant="contained">
										<ListItemButton>
											{t("edit")}
										</ListItemButton>
										<ListItemButton
											onClick={() =>
												handleDeleteCompany(company.id)
											}>
											{t("delete")}
										</ListItemButton>
									</ButtonGroup>
								</ListItem>
								{i < companies.length - 1 && <Divider />}
							</>
						);
					})}
				</List>
			</Paper>
			<Button variant="contained" className="py-3">
				{t("addCompany")}
			</Button>
		</>
	);
};

export default CompaniesPage;
