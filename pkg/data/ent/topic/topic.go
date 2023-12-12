// Code generated by ent, DO NOT EDIT.

package topic

const (
	// Label holds the string label denoting the topic type in the database.
	Label = "topic"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTopicName holds the string denoting the topic_name field in the database.
	FieldTopicName = "topic_name"
	// EdgeAnalysis holds the string denoting the analysis edge name in mutations.
	EdgeAnalysis = "analysis"
	// EdgeRatings holds the string denoting the ratings edge name in mutations.
	EdgeRatings = "ratings"
	// Table holds the table name of the topic in the database.
	Table = "topics"
	// AnalysisTable is the table that holds the analysis relation/edge.
	AnalysisTable = "analyses"
	// AnalysisInverseTable is the table name for the Analysis entity.
	// It exists in this package in order to avoid circular dependency with the "analysis" package.
	AnalysisInverseTable = "analyses"
	// AnalysisColumn is the table column denoting the analysis relation/edge.
	AnalysisColumn = "topic_analysis"
	// RatingsTable is the table that holds the ratings relation/edge.
	RatingsTable = "ratings"
	// RatingsInverseTable is the table name for the Rating entity.
	// It exists in this package in order to avoid circular dependency with the "rating" package.
	RatingsInverseTable = "ratings"
	// RatingsColumn is the table column denoting the ratings relation/edge.
	RatingsColumn = "topic_ratings"
)

// Columns holds all SQL columns for topic fields.
var Columns = []string{
	FieldID,
	FieldTopicName,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}
