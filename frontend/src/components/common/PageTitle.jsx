import PropTypes from "prop-types";

function PageTitle({ title, description }) {
  return (
    <section className="page-title">
      <h2>{title}</h2>
      {description ? <p>{description}</p> : null}
    </section>
  );
}

PageTitle.propTypes = {
  title: PropTypes.string.isRequired,
  description: PropTypes.string,
};

PageTitle.defaultProps = {
  description: "",
};

export default PageTitle;
