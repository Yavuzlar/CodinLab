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
  IconButton,
  TextField,
  Button
} from "@mui/material";
import { useTheme } from "@mui/material/styles";
import Image from "next/image";
import userIcon from "../assets/icons/icons8-male-user-100.png";
import awardIcon from "../assets/icons/icons8-award-100.png";
import tupeIcon from "../assets/icons/icons8-test-tube-100.png";
import FilterUser from "../components/filter/FilterUser";
import Translations from "src/components/Translations";
import { useDispatch, useSelector } from "react-redux";
import { getAdminUser } from "src/store/user/userSlice";
import LanguageIcon from "src/components/language-icon/LanguageIcon";
import EditIcon from "../assets/icons/icons8-edit-64.png";
import DeleteIcon from "../assets/icons/icons8-delete-30.png";
import Dialog from "@mui/material/Dialog";
import DialogActions from "@mui/material/DialogActions";
import DialogContent from "@mui/material/DialogContent";
import DialogContentText from "@mui/material/DialogContentText";
import DialogTitle from "@mui/material/DialogTitle";
import MenuItem from "@mui/material/MenuItem";
import FormControl from "@mui/material/FormControl";
import Select from "@mui/material/Select";
import { deleteUserById, fetchUserById, updateUserById } from "src/store/admin/adminSlice";
import { hexToRGBA } from "src/utils/hex-to-rgba";

const UsersList = () => {
  const theme = useTheme();
  const dispatch = useDispatch();
  const { user: stateUser } = useSelector((state) => state);

  const [openDialog, setOpenDialog] = useState(false);
  const [selectedUser, setSelectedUser] = useState(null);
  const [editData, setEditData] = useState({
    githubProfile: "",
    name: "",
    role: "",
    surname: "",
    username: "",
  });

  const [filters, setFilters] = useState({
    status: "all",
    search: "",
    sort: "",
  });

  useEffect(() => {
    dispatch(getAdminUser());
  }, [dispatch]);

  const handleEdit = (user) => {
    dispatch(fetchUserById(user.userID)).then((response) => {
      const fetchedUser = response.payload;

      setEditData({
        githubProfile: fetchedUser?.githubProfile,
        name: fetchedUser?.name,
        surname: fetchedUser?.surname,
        role: fetchedUser?.role,
        username: fetchedUser?.username,
      });

      setSelectedUser(user);
      setOpenDialog(true);
    });
  };

  const handleSave = () => {
    dispatch(updateUserById({ data: editData, userid: selectedUser.userID }));
    setOpenDialog(false);
  };

  const handleDelete = (id) => {
    dispatch(deleteUserById(id));
  };

  const inputStyle = {
    borderRadius: "15px",
    border: "2px solid #0A3B7A",

    "& .MuiInputBase-root": {
      backgroundColor: (theme) =>
        `${theme.palette.background.default} !important`,
    },

    "& .MuiOutlinedInput-root": {
      borderRadius: "15px",
      "&.Mui-focused fieldset": {
        borderColor: (theme) => `${theme.palette.primary.dark} !important`,
      },
    },

    "& .MuiInputBase-input": {
      color: "#0A3B7A",
      fontWeight: "bold",
      marginTop: "5px",
    },

    "& .MuiInputLabel-root": {
      color: (theme) => `${theme.palette.primary.dark} !important`,
      fontWeight: "bold",
    },

    "& .Mui-focused": {
      "& .MuiOutlinedInput-notchedOutline": {
        borderColor: "#0A3B7A",
      },
    },
  };

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
                      gap: "0.5rem",
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
                <TableCell
                  sx={{
                    borderBottom: "none",
                    fontFamily: "Outfit",
                    fontSize: "1rem",
                    lineHeight: "normal",
                    padding: "10px 10px 10px 0px",
                    whiteSpace: "nowrap",
                    width: "10%",
                  }}
                  align="left"
                >
                  <Translations text="userlist.action.name" />
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
                  <TableCell
                    align="left"
                    sx={{
                      borderBottom: "none",
                      padding: "10px 10px 10px 0px",
                    }}
                  >
                    <IconButton
                      onClick={() => handleEdit(row)}
                      aria-label="edit"
                    >
                      <Image
                        src={EditIcon}
                        alt="Edit Icon"
                        width={25}
                        height={25}
                        style={{ filter: 'invert(100%)' }} 
                      />
                    </IconButton>
                    <IconButton
                      onClick={() => handleDelete(row.userID)}
                      aria-label="delete"
                    >
                      <Image
                        src={DeleteIcon}
                        alt="Delete Icon"
                        width={25}
                        height={25}
                        style={{ filter: 'invert(100%)' }} 
                      />
                    </IconButton>
                  </TableCell>
                </TableRow>
              ))}
            </TableBody>
          </Table>
        </TableContainer>
      </Grid>

      <Dialog open={openDialog} onClose={() => setOpenDialog(false)}>
        <DialogTitle>
          <Translations text="userlist.edit.desc" />
        </DialogTitle>
        <DialogContent
          sx={{
            color: (theme) => `${theme.palette.text.primary} !important`,
          }}
        >
          <DialogContentText
            sx={{
              color: hexToRGBA(theme.palette.text.primary, 0.7),
            }}
          >
            <Translations text="userlist.edit.content"/>
          </DialogContentText>
          <TextField
            autoFocus
            margin="dense"
            label={<Translations text="register.name"/>}
            type="text"
            fullWidth
            variant="filled"
            value={editData.name}
            onChange={(e) => setEditData({ ...editData, name: e.target.value })}
            sx={{ ...inputStyle }}
          />
          <TextField
            margin="dense"
            label={<Translations text="register.surname"/>}
            type="text"
            fullWidth
            variant="filled"
            value={editData.surname}
            sx={{ ...inputStyle }}
            onChange={(e) =>
              setEditData({ ...editData, surname: e.target.value })
            }
          />
          <TextField
            margin="dense"
            label={<Translations text="register.username"/>}
            type="text"
            fullWidth
            variant="filled"
            value={editData.username}
            sx={{ ...inputStyle }}
            onChange={(e) =>
              setEditData({ ...editData, username: e.target.value })
            }
          />
          <TextField
            margin="dense"
            label="Level"
            type="text"
            fullWidth
            variant="filled"
            value={editData.githubProfile}
            sx={{ ...inputStyle }}
            onChange={(e) =>
              setEditData({ ...editData, githubProfile: e.target.value })
            }
          />
          <FormControl fullWidth>
            <Select
              labelId="demo-simple-select-label"
              id="demo-simple-select"
              displayEmpty
              value={editData.role}
              onChange={(e) =>
                setEditData({ ...editData, role: e.target.value })
              }
              sx={{
                ...inputStyle,
                marginTop: "10px",
                "& .MuiSelect-select": {
                  color: "#0A3B7A",
                  fontWeight: "bold",
                },
                "& .MuiSelect-icon": {
                  color: "#0A3B7A",
                },
                backgroundColor: (theme) =>
                  `${theme.palette.background.default} !important`,
              }}
            >
              <MenuItem value={"admin"}>Admin</MenuItem>
              <MenuItem value={"user"}>User</MenuItem>
            </Select>
          </FormControl>
        </DialogContent>
        <DialogActions
          sx={{
            display: "flex",
            justifyContent: "space-between",
            padding: "24px",
          }}
        >
          <Button onClick={() => setOpenDialog(false)} variant="dark">
            <Translations text="dialog.button.cancel" />
          </Button>
          <Button onClick={handleSave} color="primary" variant="dark">
            <Translations text="dialog.button.submit" />
          </Button>
        </DialogActions>
      </Dialog>
    </Grid>
  );
};

export default UsersList;
