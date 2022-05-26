import React from 'react';
import { Link } from 'react-router-dom';

import { EventModel } from '../models/event';

type Props = {
  event: EventModel;
}

export const EventItem = ({ event }: Props) => {

  return (
    <div>
      <div>
        <span>{event.name}</span>
      </div>
      <div>
        <Link to={`/events/${event.id}`}>Show</Link>
      </div>
    </div>
  )
};
