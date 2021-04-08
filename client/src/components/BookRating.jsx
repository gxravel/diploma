import React, { useEffect, useState } from 'react';
import { withStyles } from '@material-ui/core/styles';
import Rating from '@material-ui/lab/Rating';
import FavoriteIcon from '@material-ui/icons/Favorite';
import Typography from '@material-ui/core/Typography';
import Box from '@material-ui/core/Box';
import { addRating, getRating, editRating } from '../requests/rating';

const StyledRating = withStyles({
  iconFilled: {
    color: '#ff6d75',
  },
  iconHover: {
    color: '#ff3d47',
  },
})(Rating);

export default function BookRating(props) {
  const [userValue, setUserValue] = useState(0);
  const [rating, setRating] = useState({
    avg_value: 0,
    number: 0,
  });

  const { bookId, token } = props;

  const handleClick = value => {
    let method;
    if (userValue === 0) {
      method = addRating;
      setRating(prevState => {
        return {
          ...prevState,
          number: prevState.number + 1,
        };
      });
    } else {
      method = editRating;
    }
    method(value, bookId, token).then(() => {
      setUserValue(value);
    });
  };

  useEffect(() => {
    getRating(bookId, token).then(rating => {
      setRating(rating);
      setUserValue(rating.user_value);
    });
  }, [bookId, token]);

  return (
    <div
      style={{
        marginTop: '16px',
        display: 'flex',
        flexDirection: 'column',
        justifyContent: 'center',
      }}
    >
      <Box component="fieldset" borderColor="transparent">
        <Typography component="legend">Оценить книгу</Typography>
        <StyledRating
          name="book-rating"
          value={userValue}
          precision={1}
          onChange={e => handleClick(parseInt(e.currentTarget.value))}
          icon={<FavoriteIcon fontSize="inherit" />}
        />
      </Box>
      {rating && (
        <Box component="fieldset" borderColor="transparent">
          <Typography component="legend">Рейтинг</Typography>
          {rating.avg_value ? (
            <Typography>{`${rating.avg_value} (${rating.number})`}</Typography>
          ) : (
            <Typography>{`- (${rating.number})`}</Typography>
          )}
        </Box>
      )}
    </div>
  );
}
