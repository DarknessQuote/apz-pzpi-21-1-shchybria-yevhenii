import { RouterProvider, createBrowserRouter } from "react-router-dom";

function App() {
	const router = createBrowserRouter([
		{
			path: "/",
			element: <div>Main Element</div>,
		},
	]);

	return <RouterProvider router={router} />;
}

export default App;
