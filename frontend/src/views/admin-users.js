import { useEffect, useState } from "react";
import {
  Box,
  Grid,
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
  Paper,
} from "@mui/material";
import { useTheme } from "@mui/material/styles";
import Image from "next/image";
import userIcon from "../assets/icons/icons8-male-user-100.png";
import awardIcon from "../assets/icons/icons8-award-100.png";
import tupeIcon from "../assets/icons/icons8-test-tube-100.png";
import FilterUser from "../components/filter/FilterUser";
import Translations from "src/components/Translations";
import i18n from "src/configs/i18n";
import { useDispatch, useSelector } from "react-redux";
import { getAdminUser } from "src/store/user/userSlice";
import { Language } from "@mui/icons-material";
import LanguageIcon from "src/components/language-icon/LanguageIcon";

const UsersList = () => {
  const theme = useTheme();
  const dispatch = useDispatch();
  const { user: stateUser } = useSelector((state) => state);

  useEffect(() => {
    dispatch(getAdminUser());
  }, [dispatch]);

  const [filters, setFilters] = useState({
    status: "all",
    search: "",
    sort: "",
  });

  const language = i18n.language;

  return (
    <Grid container spacing={2} direction="column">
      <Grid item xs={12}>
        <Box
          sx={{
            display: "flex",
            gap: "1rem",
            flexDirection: "column",
            height: "100%",
          }}
        >
          <FilterUser filters={filters} setFilters={setFilters} />
        </Box>
      </Grid>
      <Grid
        item
        xs={12}
        sx={{
          maxHeight: "calc(100vh - 143px)",
          overflow: "auto",
        }}
      >
        <TableContainer
          component={Paper}
          sx={{ borderRadius: "15px 15px 0px 0px" }}
        >
          <Table sx={{ minWidth: 650 }} aria-label="simple table">
            <TableHead
              sx={{
                bgcolor: theme.palette.primary.dark,
              }}
            >
              <TableRow>
                <TableCell
                  sx={{
                    borderBottom: "none",
                    fontFamily: "Outfit",
                    fontSize: "1rem",
                    lineHeight: "normal",
                    padding: "10px 10px 10px 25px",
                    whiteSpace: "nowrap",
                    width: "20%",
                  }}
                >
                  <Translations text="userlist.order.name" />
                </TableCell>
                <TableCell
                  sx={{
                    display: "flex",
                    alignItems: "start",
                    justifyContent: "start",
                    borderBottom: "none",
                    fontFamily: "Outfit",
                    fontSize: "1rem",
                    lineHeight: "normal",
                    padding: "10px 10px 10px 0px",
                    whiteSpace: "nowrap",
                    width: "20%",
                  }}
                >
                  <Box
                    sx={{
                      display: "flex",
                      alignItems: "center",
                      justifyContent: "center",
                      gap: "0.5rem",
                      padding: 0,
                      margin: 0,
                    }}
                  >
                    <Image
                      src={userIcon}
                      alt="User Icon"
                      width={25}
                      height={25}
                    />
                    <Box>
                      <Translations text="userlist.username.name" />
                    </Box>
                  </Box>
                </TableCell>
                <TableCell
                  sx={{
                    borderBottom: "none",
                    fontFamily: "Outfit",
                    fontSize: "1rem",
                    lineHeight: "normal",
                    padding: "10px 10px 10px 0px",
                    whiteSpace: "nowrap",
                    width: "25%",
                  }}
                >
                  <Box
                    sx={{
                      display: "flex",
                      alignItems: "center",
                      justifyContent: "start",
                      gap: "0.5rem",
                      padding: 0,
                      margin: 0,
                    }}
                  >
                    <Image
                      src={tupeIcon}
                      alt="Tupe Icon"
                      width={25}
                      height={25}
                    />
                    <Translations text="userlist.level.name" />
                  </Box>
                </TableCell>
                <TableCell
                  sx={{
                    borderBottom: "none",
                    fontFamily: "Outfit",
                    fontSize: "1rem",
                    lineHeight: "normal",
                    padding: "10px 10px 10px 0px",
                    whiteSpace: "nowrap",
                    width: "25%",
                  }}
                >
                  <Box
                    sx={{
                      display: "flex",
                      alignItems: "center",
                      justifyContent: "start",
                      gap: "0.5rem",
                      padding: 0,
                      margin: 0,
                    }}
                  >
                    <Image
                      src={awardIcon}
                      alt="Award Icon"
                      width={25}
                      height={25}
                    />
                    <Translations text="userlist.best.name" />
                  </Box>
                </TableCell>
              </TableRow>
            </TableHead>
            <TableBody>
              {stateUser.adminUserData?.data?.map((row) => (
                <TableRow
                  key={row.order}
                  sx={{
                    bgcolor:
                      row.order % 2 === 1
                        ? theme.palette.primary.main
                        : theme.palette.primary.light,
                  }}
                >
                  <TableCell
                    sx={{
                      fontFamily: "Outfit",
                      fontSize: "1rem",
                      lineHeight: "normal",
                      padding: "10px 10px 10px 25px",
                      borderBottom: "none",
                    }}
                  >
                    {row.order}
                  </TableCell>
                  <TableCell
                    sx={{
                      borderBottom: "none",
                      fontFamily: "Outfit",
                      fontSize: "1rem",
                      lineHeight: "normal",
                      padding: "10px 10px 10px 35px",
                    }}
                    align="left"
                  >
                    {row.username}
                  </TableCell>
                  <TableCell
                    sx={{
                      borderBottom: "none",
                      fontFamily: "Outfit",
                      fontSize: "1rem",
                      lineHeight: "normal",
                      padding: "10px 10px 10px 35px",
                    }}
                    align="left"
                  >
                    {row.level}
                  </TableCell>
                  <TableCell
                    align="left"
                    sx={{
                      borderBottom: "none",
                      fontFamily: "Outfit",
                      fontSize: "1rem",
                      lineHeight: "normal",
                      padding: "10px 10px 10px 35px",
                    }}
                  >
                    <LanguageIcon language={row.bestLanguage} />
                  </TableCell>
                </TableRow>
              ))}
            </TableBody>
          </Table>
        </TableContainer>
      </Grid>
    </Grid>
  );
};

export default UsersList;
