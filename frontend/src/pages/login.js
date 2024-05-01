const { default: BlankLayout } = require("@/layout/BlankLayout")

const Login = () => {
    return (
        <div>login</div>
    )
}


Login.guestGuard = true
Login.getLayout = page => <BlankLayout>{page}</BlankLayout>

export default Login