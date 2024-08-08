import { Box } from "@mui/material";
import SkeletonLoader from "src/components/skeleton/SkeletonLoader.js";

const Sample = () => {
  const data = [ // this is the data that will be passed to the SkeletonLoader component
    { width: 350, height: 150, variant: "rectangular", animation: "wave" },
    { width: 40, height: 40, variant: "circular", animation: "pulse" },
    { width: "100%", height: "20px", variant: "text", animation: false },
  ];

  return (
    <div>
      <Box
        sx={{
          display: "flex",
          gap: "20px",
          flexWrap: "wrap",
        }}
      >
        <SkeletonLoader items={data} /> 
        <SkeletonLoader items={data} />
        <SkeletonLoader items={data} />
        <SkeletonLoader items={data} />
        <SkeletonLoader items={data} />
        <SkeletonLoader items={data} />
      </Box>
    </div>
  );
};

export default Sample;
