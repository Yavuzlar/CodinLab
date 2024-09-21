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
  Button,
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
import InputLabel from "@mui/material/InputLabel";
import MenuItem from "@mui/material/MenuItem";
import FormControl from "@mui/material/FormControl";
import Select from "@mui/material/Select";
import { fetchUserById, updateUserById } from "src/store/admin/adminSlice";

const UsersList = () => {
  const theme = useTheme();
  const dispatch = useDispatch();
  const { user: stateUser} = useSelector((state) => state);

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
        githubProfile: fetchedUser.githubProfile,
        name: fetchedUser.name,
        surname: fetchedUser.surname,
        role: fetchedUser.role,
        username: fetchedUser.username,
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
    console.log("Delete user with id: ", id);
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
                      justifyContent: "center",
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
                  align="right"
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
                    align="right"
                    sx={{
                      borderBottom: "none",
                      padding: "10px 10px 10px 0px",
                    }}
                  >
                    <IconButton
                      onClick={() => handleEdit(row)}
                      aria-label="edit"
                      sx={{ color: theme.palette.info.main }}
                    >
                      <Image
                        src={EditIcon}
                        alt="Edit Icon"
                        width={20}
                        height={20}
                      />
                    </IconButton>
                    <IconButton
                      onClick={() => handleDelete(row.id)}
                      aria-label="delete"
                      sx={{ color: theme.palette.error.main }}
                    >
                      <Image
                        src={DeleteIcon}
                        alt="Delete Icon"
                        width={20}
                        height={20}
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
        <DialogContent>
          <DialogContentText>
            <Translations text="userlist.edit.content" />
          </DialogContentText>
          <TextField
            autoFocus
            margin="dense"
            label="Name"
            type="text"
            fullWidth
            variant="outlined"
            value={editData.name}
            onChange={(e) => setEditData({ ...editData, name: e.target.value })}
          />
          <TextField
            margin="dense"
            label="Surname"
            type="text"
            fullWidth
            variant="outlined"
            value={editData.surname}
            onChange={(e) =>
              setEditData({ ...editData, surname: e.target.value })
            }
          />
          <TextField
            margin="dense"
            label="Username"
            type="text"
            fullWidth
            variant="outlined"
            value={editData.username}
            onChange={(e) =>
              setEditData({ ...editData, username: e.target.value })
            }
          />
          <TextField
            margin="dense"
            label="Level"
            type="text"
            fullWidth
            variant="outlined"
            value={editData.githubProfile}
            onChange={(e) =>
              setEditData({ ...editData, githubProfile: e.target.value })
            }
          />
          <FormControl fullWidth>
            <InputLabel id="demo-simple-select-label">Role</InputLabel>
            <Select
              labelId="demo-simple-select-label"
              id="demo-simple-select"
              value={editData.role}
              label="Role"
              onChange={(e) =>
                setEditData({ ...editData, role: e.target.value })
              }
            >
              <MenuItem value={"admin"}>Admin</MenuItem>
              <MenuItem value={"user"}>User</MenuItem>
            </Select>
          </FormControl>
        </DialogContent>
        <DialogActions>
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
