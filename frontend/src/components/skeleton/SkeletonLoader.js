import Skeleton from '@mui/material/Skeleton';
import Stack from '@mui/material/Stack';
import { useTheme } from '@mui/material/styles';


const SkeletonLoader = ({ items }) => {
  const theme = useTheme();


  return (
    <Stack spacing={1}>
      {items.map((item, index) => (
        <Skeleton 
          key={index}  // this is the key prop
          variant={item.variant}  // this is the variant (rectangular, circular, text)
          width={item.width} // this is the width prop
          height={item.height}  // this is the height prop
          animation={item.animation}  // this is the animation prop (wave, pulse, false)

          sx={{
            bgcolor: theme.palette.action.disabled,
          }}
        />
      ))}
    </Stack>
  );
};

export default SkeletonLoader;
