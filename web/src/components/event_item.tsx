import React from 'react';
import { Link } from 'react-router-dom';

import { EventModel } from '../models/event';

type Props = {
  event: EventModel;
}

export const EventItem = ({ event }: Props) => {

  return (
    <div className='event-item'>
      <div className='event-item__text'>
        <span>{event.name}</span>
      </div>
      <div className='event-item__action'>
        <Link to={`/events/${event.id}`}>Show</Link>
      </div>
    </div>
  )
};
