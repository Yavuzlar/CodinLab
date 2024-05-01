
import Image from "next/image"

const Logo = () => {
    return (
        <Image
            width={200}
            height={200}
            src={"/images/logo.png"}
            alt="logo"
            style={{
                width: 'auto',
                height: '4.5rem',
                cursor: 'pointer'
            }}
        />
    )
}

export default Logo