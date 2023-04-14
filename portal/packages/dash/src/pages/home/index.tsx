import {Layout} from 'antd';

const {Content} = Layout;

export function Home() {
	return (
		<Layout
			className="flex justify-center items-center"
			style={{
				height: '100vh',
			}}
		>
			<Content className="bg-white rounded shadow p-8 w-96 max-h-72">
				<h1 className="mb-4 text-lg text-black">Home</h1>
			</Content>
		</Layout>
	);
}