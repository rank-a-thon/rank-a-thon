import React from 'react';
import { render } from '@testing-library/react';

import { PagePublicHome } from '..';

describe('<PagePublicHome  />', () => {
  it('should match snapshot', () => {
    const loadingIndicator = render(<PagePublicHome />);
    expect(loadingIndicator.container.firstChild).toMatchSnapshot();
  });
});
