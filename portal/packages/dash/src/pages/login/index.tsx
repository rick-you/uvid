import {Form, Input, Button, message, Layout} from 'antd';
import {UserOutlined, LockOutlined} from '@ant-design/icons';
import {login} from '../../lib/request';

const {Content} = Layout;

export function Login() {
	return (
		<Layout
			className="flex justify-center items-center"
			style={{
				height: '100vh',
			}}
		>
			<Content className="bg-white rounded shadow p-8 w-96 max-h-72">
				<h1 className="mb-4 text-lg text-black">Login</h1>
				<LoginForm />
			</Content>
		</Layout>
	);
}

type FormValues = {
	name: string;
	password: string;
};

function LoginForm() {
	const onFinish = async (values: FormValues) => {
		try {
			await login(values.name, values.password);
			await message.success('Login success!');
		} catch (error) {
			console.error(error);
			await message.error('Login failed!');
		}
	};

	return (
		<div className="form-wrapper">
			<Form name="normal_login" className="login-form" onFinish={onFinish}>
				<Form.Item
					name="name"
					rules={[
						{
							required: true,
						},
					]}
				>
					<Input
						prefix={<UserOutlined className="site-form-item-icon" />}
						placeholder="Username"
					/>
				</Form.Item>
				<Form.Item
					name="password"
					rules={[
						{
							required: true,
						},
					]}
				>
					<Input.Password
						prefix={<LockOutlined className="site-form-item-icon" />}
						type="password"
						placeholder="Password"
					/>
				</Form.Item>

				<Form.Item>
					<Button
						type="primary"
						htmlType="submit"
						className="login-form-button"
					>
						Log in
					</Button>
				</Form.Item>
			</Form>
		</div>
	);
}