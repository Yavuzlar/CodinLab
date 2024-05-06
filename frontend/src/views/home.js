import { Button } from '@mui/material';
import { useTheme } from '@mui/material/styles';



const Home = () => {
    const theme = useTheme();
    return (
        <div><Button sx={{ ...theme.custombutton.lightButton}}> TEST </Button></div>
    )
}

export default Home