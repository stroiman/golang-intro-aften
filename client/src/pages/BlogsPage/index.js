import React from 'react';
import { connect } from 'react-redux';
import PropTypes from 'prop-types';
import * as getters from '../../reducers';
import * as actions from './actions';
import PageHeader from '../../components/PageHeader';

class BlogsPage extends React.Component {
  constructor(props){
    super(props);
    this.state = {
      loadingState: props.loadingState
    };
  };

  componentWillReceiveProps(newProps) {
    const { loadingState } = newProps;
    this.setState({ loadingState });
  }

  componentDidMount() {
    this.props.loadBlogs();
  };

  render() {
    return (<div>
      <PageHeader>Find one blog that interests you</PageHeader>
      { this.state.loadingState }
    </div>);
  }
}

BlogsPage.propTypes = {
  loadingState: PropTypes.string.isRequired
};

const mapStateToProps = state => ({
  loadingState: getters.getBlogsLoadingState(state)
})

export default connect(mapStateToProps, actions)(BlogsPage);
