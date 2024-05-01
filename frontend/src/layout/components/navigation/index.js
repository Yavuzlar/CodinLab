import List from '@mui/material/List';
import navigation from '@/navigation';
import NavItem from './item/NavItem';

const NavigationList = () => {
    return (
        <List
            sx={{
                display: 'flex',
                gap: '0.5rem',
                flexDirection: 'column',
                height: '100%',
                pt: 0
            }}
            component="nav"
        >
            {
                navigation.map((item, index) => <NavItem key={index} {...item} />)
            }
        </List >
    );
}

export default NavigationList;