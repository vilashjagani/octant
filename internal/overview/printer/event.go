package printer

import (
	"fmt"
	"strings"

	"github.com/heptio/developer-dash/internal/view/flexlayout"

	"github.com/pkg/errors"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/runtime"

	"github.com/heptio/developer-dash/internal/cache"
	"github.com/heptio/developer-dash/internal/view/component"
	"github.com/heptio/developer-dash/internal/view/gridlayout"
)

var (
	objectEventCols = component.NewTableCols("Message", "Reason", "Type", "First Seen", "Last Seen", "From", "Count")
)

// EventListHandler is a printFunc that lists events.
func EventListHandler(list *corev1.EventList, opts Options) (component.ViewComponent, error) {
	if list == nil {
		return nil, errors.New("nil list")
	}

	cols := component.NewTableCols("Kind", "Message", "Reason", "Type",
		"First Seen", "Last Seen")
	table := component.NewTable("Events", cols)

	for _, event := range list.Items {
		row := component.TableRow{}

		objectPath, err := ObjectReferencePath(event.InvolvedObject)
		if err != nil {
			return nil, err
		}

		infoItems := []component.ViewComponent{
			component.NewLink("", event.InvolvedObject.Name, objectPath),
			component.NewText(fmt.Sprintf("%d", event.Count)),
		}
		info := component.NewList("", infoItems)

		row["Kind"] = info
		eventPath := gvkPath(event.Namespace, event.TypeMeta.APIVersion, event.TypeMeta.Kind, event.Name)
		row["Message"] = component.NewLink("", event.Message, eventPath)
		row["Reason"] = component.NewText(event.Reason)
		row["Type"] = component.NewText(event.Type)
		row["First Seen"] = component.NewTimestamp(event.FirstTimestamp.Time)
		row["Last Seen"] = component.NewTimestamp(event.LastTimestamp.Time)

		table.Add(row)
	}

	return table, nil
}

func EventHandler(event *corev1.Event, opts Options) (component.ViewComponent, error) {
	if event == nil {
		return nil, errors.New("event can not be nil")
	}

	var detailSections []component.SummarySection

	detailSections = append(detailSections, component.SummarySection{
		Header:  "Last Seen",
		Content: component.NewTimestamp(event.LastTimestamp.Time),
	})

	detailSections = append(detailSections, component.SummarySection{
		Header:  "First Seen",
		Content: component.NewTimestamp(event.FirstTimestamp.Time),
	})

	detailSections = append(detailSections, component.SummarySection{
		Header:  "Count",
		Content: component.NewText(fmt.Sprintf("%d", event.Count)),
	})

	detailSections = append(detailSections, component.SummarySection{
		Header:  "Message",
		Content: component.NewText(event.Message),
	})

	detailSections = append(detailSections, component.SummarySection{
		Header:  "Kind",
		Content: component.NewText(event.InvolvedObject.Kind),
	})

	// NOTE: object reference can contain a field path to the object,
	// and that is not reported.
	refPath, err := ObjectReferencePath(event.InvolvedObject)
	if err != nil {
		return nil, err
	}
	detailSections = append(detailSections, component.SummarySection{
		Header:  "Involved Object",
		Content: component.NewLink("", event.InvolvedObject.Name, refPath),
	})

	detailSections = append(detailSections, component.SummarySection{
		Header:  "Type",
		Content: component.NewText(event.Type),
	})

	detailSections = append(detailSections, component.SummarySection{
		Header:  "Reason",
		Content: component.NewText(event.Reason),
	})

	sourceMsg := event.Source.Component
	if event.Source.Host != "" {
		sourceMsg = fmt.Sprintf("%s on %s",
			event.Source.Component, event.Source.Host)
	}
	detailSections = append(detailSections, component.SummarySection{
		Header:  "Source",
		Content: component.NewText(sourceMsg),
	})

	summary := component.NewSummary("Event Detail", detailSections...)

	gl := gridlayout.New()

	summarySection := gl.CreateSection(10)
	summarySection.Add(summary, 24)

	grid := gl.ToGrid()
	return grid, nil
}

// PrintEvents collects events for a resource
func PrintEvents(list *corev1.EventList, opts Options) (component.ViewComponent, error) {
	if list == nil {
		return nil, errors.New("nil list")
	}

	table := component.NewTable("Events", objectEventCols)

	for _, event := range list.Items {
		row := component.TableRow{}

		row["Message"] = component.NewText(event.Message)
		row["Reason"] = component.NewText(event.Reason)
		row["Type"] = component.NewText(event.Type)

		row["First Seen"] = component.NewTimestamp(event.FirstTimestamp.Time)
		row["Last Seen"] = component.NewTimestamp(event.LastTimestamp.Time)

		row["From"] = component.NewText(formatEventSource(event.Source))

		count := fmt.Sprintf("%d", event.Count)
		row["Count"] = component.NewText(count)

		table.Add(row)
	}

	return table, nil
}

// formatEventSource formats EventSource as a comma separated string excluding Host when empty
func formatEventSource(es corev1.EventSource) string {
	EventSourceString := []string{es.Component}
	if len(es.Host) > 0 {
		EventSourceString = append(EventSourceString, es.Host)
	}
	return strings.Join(EventSourceString, ", ")
}

func createEventsForObject(fl *flexlayout.FlexLayout, object runtime.Object, opts Options) error {
	eventList, err := eventsForObject(object, opts.Cache)
	if err != nil {
		return errors.Wrap(err, "list events for object")
	}

	if len(eventList.Items) > 0 {
		eventTable, err := PrintEvents(eventList, opts)
		if err != nil {
			return errors.Wrap(err, "create event table for object")
		}

		eventsSection := fl.AddSection()
		if err := eventsSection.Add(eventTable, 24); err != nil {
			return errors.Wrap(err, "add event table to layout")
		}
	}

	return nil
}

func eventsForObject(object runtime.Object, c cache.Cache) (*corev1.EventList, error) {
	accessor := meta.NewAccessor()

	namespace, err := accessor.Namespace(object)
	if err != nil {
		return nil, errors.Wrap(err, "get namespace for object")
	}

	apiVersion, err := accessor.APIVersion(object)
	if err != nil {
		return nil, errors.Wrap(err, "Get apiVersion for object")
	}

	kind, err := accessor.Kind(object)
	if err != nil {
		return nil, errors.Wrap(err, "get kind for object")
	}

	name, err := accessor.Name(object)
	if err != nil {
		return nil, errors.Wrap(err, "get name for object")
	}

	key := cache.Key{
		Namespace:  namespace,
		APIVersion: "v1",
		Kind:       "Event",
	}

	list, err := c.List(key)
	if err != nil {
		return nil, errors.Wrap(err, "list events for object")
	}

	eventList := &corev1.EventList{}

	for _, unstructuredEvent := range list {
		event := &corev1.Event{}
		err := runtime.DefaultUnstructuredConverter.FromUnstructured(unstructuredEvent.Object, event)
		if err != nil {
			return nil, err
		}

		involvedObject := event.InvolvedObject
		if involvedObject.Namespace == namespace &&
			involvedObject.APIVersion == apiVersion &&
			involvedObject.Kind == kind &&
			involvedObject.Name == name {
			eventList.Items = append(eventList.Items, *event)
		}
	}

	return eventList, nil
}