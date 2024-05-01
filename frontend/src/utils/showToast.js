import toast from 'react-hot-toast'

// return rejectWithValue(error.response.data.code);

const showToast = (type, message, position, style) => {
  let duration = 2000
  if (message?.length > 15) {
    duration += (message?.length - 15) * 25
  }

  if (message && message != '') {
    toast[type]?.(message, {
      position: position ?? 'top-center',
      duration: duration,
      style: style,
    });
  } else {
    if (type == 'dismiss') toast[type]()
  }
}

export { showToast }