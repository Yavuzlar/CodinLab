import { Box, Divider } from "@mui/material";

const Output = ({ value, params }) => {
    value = value || "Output will be displayed here";

   
    const adjustedHeight = params && params.height ? `calc(${params.height} + 6.7vh)` : "82vh";

    return (
        <Box
            sx={{
                display: "flex",
                flexDirection: "column",
                gap: "10px",
                width: "50%",
                height: adjustedHeight,
                padding: "10px",
                borderRadius: "30px",
                opacity: "1",
                backgroundColor: "black", 
                color: "white",
            }}
        >
            <Box
                sx={{
                    display: "flex",
                    justifyContent: "space-between",
                    fontSize: "15px",
                    fontWeight: "bold",
                    alignItems: "center",
                    paddingLeft: "10px",
                    paddingRight: "10px",
                }}
            >
                <div
                    sx={{
                        display: "flex",
                    }}
                >
                    <div
                        style={{
                            display: "flex",
                            gap: "10px",
                        }}
                    >
                        Output
                    </div>
                </div>
            </Box>
            <Divider
                sx={{
                    width: "103%",
                    height: "2px",
                    marginLeft: "-10px",
                    backgroundColor: "white",
                }}
            />
            <Box>
                <div
                    style={{
                        padding: "10px",
                        fontSize: "20px",
                        width: "100%",
                        height: "100%",
                    }}
                >
                    {value}
                </div>
            </Box>
        </Box>
    );
};

export default Output;
