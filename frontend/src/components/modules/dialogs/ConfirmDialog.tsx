import * as React from 'react';
import Button from '@mui/material/Button';
import Dialog from '@mui/material/Dialog';
import DialogActions from '@mui/material/DialogActions';
import DialogContent from '@mui/material/DialogContent';
import DialogContentText from '@mui/material/DialogContentText';
import DialogTitle from '@mui/material/DialogTitle';
import { Slide, SlideProps } from '@mui/material';
const Transition = React.forwardRef(function Transition(
  props: SlideProps & { children?: React.ReactNode }, 
  ref
) {
  return <Slide direction="up" ref={ref} {...props} />;
});

export default function ConfirmDialog(props: any) {
  const { open, text, message, callback } = props;

  const handleOK = () => {
    callback(true);
  };

  const handleClose = () => {
    callback(false);
  };

  return (
    <Dialog
      open={open}
      TransitionComponent={Transition}
      keepMounted
      onClose={handleClose}
      aria-describedby="alert-dialog-slide-description"
    >
      <DialogTitle>{"Confirm"}</DialogTitle>
      <DialogContent>
        <DialogContentText id="alert-dialog-slide-description">
          {message}
        </DialogContentText>
      </DialogContent>
      <DialogActions>
        <Button onClick={handleClose}>CANCEL</Button>
        <Button onClick={handleOK} variant="contained">OK</Button>
      </DialogActions>
    </Dialog>
  );
}
